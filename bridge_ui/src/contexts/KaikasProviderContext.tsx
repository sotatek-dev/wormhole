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
import { KLAYTN_NETWORK_CHAIN_ID } from "../utils/consts";
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
  signerAddress: string | undefined;
  providerError: string | null;
}

const KaikasProviderContext = createContext<IKaikasProviderContext>({
  connect: () => {},
  disconnect: () => {},
  provider: undefined,
  chainId: undefined,
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
  const [signerAddress, setSignerAddress] = useState<string | undefined>(
    undefined
  );

  const connect = useCallback(async () => {
    const { klaytn } = window;
    const caver = new Caver(klaytn);
    setProviderError("Waiting for unlocking kaikas wallet...");
    setProvider(caver.klay);

    if (klaytn) {
      setKlaytnApi(klaytn);
      try {
        const { result: accounts } = await klaytn.send("klay_requestAccounts", []);
        if (accounts) {
          setSignerAddress(accounts[0]);
          setProviderError(null);
        }
        else {
          setProviderError("An error occurred while getting the signer address");
        }
        setChainId(klaytn.networkVersion);
      } catch (error) {
        console.log("User denied account access");
      }
    } else {
      setProviderError("Please install Kaikas");
    }
  }, []);

  useEffect(() => {
    const mutatorNetworkChanged = async () => {
      const _chainId = await provider.getChainId();
      setChainId(_chainId);
    };

    const mutatorAccountsChanged = async () => {
      try {
        const _addresses = await provider.getAccounts();
        setSignerAddress(_addresses[0]);
      } catch (error) {
        setProviderError(
          "An error occurred while getting the signer address"
        );
      }
    }
    
    if (provider && klaytnApi && klaytnApi.on) {
      klaytnApi.on("networkChanged", mutatorNetworkChanged);
      klaytnApi.on("accountsChanged", mutatorAccountsChanged);
    }

    return () => {
      klaytnApi?.on("networkChanged", mutatorNetworkChanged);
      klaytnApi?.on("accountsChanged", mutatorAccountsChanged);
    }
  },[provider, klaytnApi])
  
  useEffect(() => {
    if (chainId !== KLAYTN_NETWORK_CHAIN_ID) {
      signerAddress && setProviderError(
        `Wallet is not connected to ${process.env.REACT_APP_CLUSTER}! Expected Chain ID: ${KLAYTN_NETWORK_CHAIN_ID}!`
      )
    }
    else {
      setProviderError(null);
    }
  }, [chainId, signerAddress])

  const disconnect = useCallback(() => {
    setProviderError(null);
    setProvider(undefined);
    setChainId(undefined);
    setSignerAddress(undefined);
  }, []);

  const contextValue = useMemo(
    () => ({
      connect,
      disconnect,
      provider,
      chainId,
      signerAddress,
      providerError,
    }),
    [
      connect,
      disconnect,
      provider,
      chainId,
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
