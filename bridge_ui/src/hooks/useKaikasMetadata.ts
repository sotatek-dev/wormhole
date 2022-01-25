import { useEffect, useMemo, useState } from "react";
import { useKaikasProvider } from "../contexts/KaikasProviderContext";
import { DataWrapper } from "../store/helpers";
import TokenImplementation from "../blockchain/abi/TokenImplementation.json";

export type KlaytnMetadata = {
  symbol?: string;
  logo?: string;
  tokenName?: string;
  decimals?: number;
};

const handleError = () => {
  return undefined;
};

const fetchSingleMetadata = async (
  address: string,
  provider: any
): Promise<KlaytnMetadata> => {
  const contract = new provider.Contract(TokenImplementation, address);

  const [name, symbol, decimals] = await Promise.all([
    contract.methods.name().call().catch(handleError),
    contract.methods.symbol().call().catch(handleError),
    contract.methods.decimals().call().catch(handleError),
  ]);
  return { tokenName: name, symbol, decimals };
};

const fetchKlaytnMetadata = async (addresses: string[], provider: any) => {
  const promises: Promise<KlaytnMetadata>[] = [];
  addresses.forEach((address) => {
    promises.push(fetchSingleMetadata(address, provider));
  });
  const resultsArray = await Promise.all(promises);
  const output = new Map<string, KlaytnMetadata>();
  addresses.forEach((address, index) => {
    output.set(address, resultsArray[index]);
  });

  return output;
};

function useKlaytnMetadata(
  addresses: string[]): DataWrapper<Map<string, KlaytnMetadata>> {
  const { provider } = useKaikasProvider();

  const [isFetching, setIsFetching] = useState(false);
  const [error, setError] = useState("");
  const [data, setData] = useState<Map<string, KlaytnMetadata> | null>(null);

  useEffect(() => {
    let cancelled = false;
    if (addresses.length && provider) {
      setIsFetching(true);
      setError("");
      setData(null);
      fetchKlaytnMetadata(addresses, provider).then(
        (results) => {
          if (!cancelled) {
            setData(results);
            setIsFetching(false);
          }
        },
        () => {
          if (!cancelled) {
            setError("Could not retrieve contract metadata");
            setIsFetching(false);
          }
        }
      );
    }
    return () => {
      cancelled = true;
    };
  }, [addresses, provider]);

  return useMemo(
    () => ({
      data,
      isFetching,
      error,
      receivedAt: null,
    }),
    [data, isFetching, error]
  );
}

export default useKlaytnMetadata;
