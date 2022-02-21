import {
  createContext,
  ReactChildren,
  useCallback,
  useContext,
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
  const [providerError, setProviderError] = useState<string | null>(null);
  const [provider, setProvider] = useState<any>(undefined);
  const [chainId, setChainId] = useState<number | undefined>(undefined);
  const [signer, setSigner] = useState<any>(undefined);
  const [signerAddress, setSignerAddress] = useState<string | undefined>(
    undefined
  );

  let isUnlocked;
  const connect = useCallback(async () => {
    const { klaytn } = window;
    setProviderError(null);
    isUnlocked=await window.klaytn._kaikas.isUnlocked();
    if(!isUnlocked){
      setProviderError("An error occurred while requesting accounts");
    }

    if (klaytn) {
      try {
        await klaytn.enable();
        const caver = new Caver(window.klaytn);
        setProvider(caver.klay);
        setSignerAddress(klaytn.selectedAddress);
        setChainId(klaytn.networkVersion);
        setProviderError(null)
        klaytn.on("networkChanged", () => setChainId(klaytn.networkVersion));
        klaytn.on("accountsChanged", () =>
          setSignerAddress(klaytn.selectedAddress)
        );
      } catch (error) {
        console.log("User denied account access");
      }
    } else {
      setProviderError("Please install Kaikas");
    }
  }, []);

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
