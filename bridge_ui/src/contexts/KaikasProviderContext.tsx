import {
  createContext,
  ReactChildren,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";
import Caver from "caver-js";
declare global {
  interface Window {
    klaytn: any;
  }
}

interface IKaikasProviderContext {
  connect(): void;
  disconnect(): void;
  provider: any;
  chainId: number | undefined;
  signer: any;
  signerAddress: string | undefined;
  providerError: string | null;
}

const KaikasProviderContext = createContext<IKaikasProviderContext>({
  connect: () => {},
  disconnect: () => {},
  provider: undefined,
  chainId: undefined,
  signer: undefined,
  signerAddress: undefined,
  providerError: null,
});

export const KaikasProviderProvider = ({
  children,
}: {
  children: ReactChildren;
  }) => {
  const [klaytnApi, setKlaytnApi] = useState<any>(null);
  const [providerError, setProviderError] = useState<string | null>(null);
  const [provider, setProvider] = useState<any>(undefined);
  const [chainId, setChainId] = useState<number | undefined>(undefined);
  const [signer, setSigner] = useState<any>(undefined);
  const [signerAddress, setSignerAddress] = useState<string | undefined>(
    undefined
  );

  const connect = useCallback(async () => {
    const { klaytn } = window;
    const caver = new Caver(klaytn);

    setProviderError(null);
    setProvider(caver.klay);

    if (klaytn) {
      setKlaytnApi(klaytn);
      try {
        const {result: accounts} = await klaytn.send("klay_requestAccounts", []);
        setSignerAddress(accounts[0]);
        setChainId(klaytn.networkVersion);
      } catch (error) {
        console.error(error);
        console.log("User denied account access");
      }
    } else {
      setProviderError("Please install Kaikas");
    }
  }, []);

  useEffect(() => {
    if (provider && klaytnApi && klaytnApi.on) {
      klaytnApi.on("networkChanged", async () => {
        console.log("klay_networkChanged");
        const _chainId = await provider.getChainId();
        setChainId(_chainId);
      });
      klaytnApi.on("accountsChanged", async () => {
        try {
          console.log("klay_accountsChanged");
          const _addresses = await provider.getAccounts();
          setSignerAddress(_addresses[0]);
        } catch (error) {
          console.error(error);
          setProviderError(
            "An error occurred while getting the signer address"
          );
        }
      });
    }
  },[provider, klaytnApi])

  const disconnect = useCallback(() => {
    setProviderError(null);
    setProvider(undefined);
    setChainId(undefined);
    setSigner(undefined);
    setSignerAddress(undefined);
  }, []);

  const contextValue = useMemo(
    () => ({
      connect,
      disconnect,
      provider,
      chainId,
      signer,
      signerAddress,
      providerError,
    }),
    [
      connect,
      disconnect,
      provider,
      chainId,
      signer,
      signerAddress,
      providerError,
    ]
  );

  return (
    <KaikasProviderContext.Provider value={contextValue}>
      {children}
    </KaikasProviderContext.Provider>
  );
};

export const useKaikasProvider = () => {
  return useContext(KaikasProviderContext);
};
