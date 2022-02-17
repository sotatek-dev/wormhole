import { useKaikasProvider } from './../contexts/KaikasProviderContext';
import {
  approveEth,
  ChainId,
  getAllowanceEth,
  isEVMChain,
  CHAIN_ID_KLAYTN_BAOBAB
} from "@certusone/wormhole-sdk";
import { BigNumber } from "ethers";
import { useEffect, useMemo, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { useEthereumProvider } from "../contexts/EthereumProviderContext";
import { selectTransferIsApproving } from "../store/selectors";
import { setIsApproving } from "../store/transferSlice";
import { getTokenBridgeAddressForChain } from "../utils/consts";
import { approveKlaytn, getAllowanceKlaytn } from '../utils/klaytn';

export default function useAllowance(
  chainId: ChainId,
  tokenAddress?: string,
  transferAmount?: BigInt,
  sourceIsNative?: boolean
) {
  const dispatch = useDispatch();
  const [allowance, setAllowance] = useState<BigInt | null>(null);
  const [isAllowanceFetching, setIsAllowanceFetching] = useState(false);
  const isApproveProcessing = useSelector(selectTransferIsApproving);
  const { signer } = useEthereumProvider();
  const { provider: providerKaikas, signerAddress: signerAddressKaikas } = useKaikasProvider();
  const sufficientAllowance =
    !(isEVMChain(chainId) || chainId === CHAIN_ID_KLAYTN_BAOBAB) ||
    sourceIsNative ||
    (allowance && transferAmount && allowance >= transferAmount);
  
  console.log({sufficientAllowance});

  useEffect(() => {
    let cancelled = false;
    if (isEVMChain(chainId) && tokenAddress && signer && !isApproveProcessing) {
      setIsAllowanceFetching(true);
      getAllowanceEth(
        getTokenBridgeAddressForChain(chainId),
        tokenAddress,
        signer
      ).then(
        (result) => {
          if (!cancelled) {
            setIsAllowanceFetching(false);
            setAllowance(result.toBigInt());
          }
        },
        (error) => {
          if (!cancelled) {
            setIsAllowanceFetching(false);
            //setError("Unable to retrieve allowance"); //TODO set an error
          }
        }
      );
    } else if (chainId === CHAIN_ID_KLAYTN_BAOBAB && tokenAddress && providerKaikas && signerAddressKaikas && !isApproveProcessing) {
      setIsAllowanceFetching(true);
      getAllowanceKlaytn(
        getTokenBridgeAddressForChain(chainId),
        tokenAddress,
        providerKaikas,
        signerAddressKaikas
      ).then(
        (result) => {
          if (!cancelled) {
            setIsAllowanceFetching(false);
            setAllowance(result);
          }
        },
        (error) => {
          if (!cancelled) {
            setIsAllowanceFetching(false);
            //setError("Unable to retrieve allowance"); //TODO set an error
          }
        }
      );
    }

    return () => {
      cancelled = true;
    };
  }, [
    chainId,
    tokenAddress,
    signer,
    isApproveProcessing,
    providerKaikas,
    signerAddressKaikas
  ]);

  const approveAmount: any = useMemo(() => {
    if (!(isEVMChain(chainId) || chainId === CHAIN_ID_KLAYTN_BAOBAB) || !tokenAddress || !signer || !providerKaikas) {
      return (amount: BigInt) => {
        return Promise.resolve();
      }
    }
    else if (isEVMChain(chainId)) {
      return (amount: BigInt) => {
        dispatch(setIsApproving(true));
        return approveEth(
          getTokenBridgeAddressForChain(chainId),
          tokenAddress,
          signer,
          BigNumber.from(amount)
        ).then(
          () => {
            dispatch(setIsApproving(false));
            return Promise.resolve();
          },
          () => {
            dispatch(setIsApproving(false));
            return Promise.reject();
          }
        );
      };
    }
    else if (chainId === CHAIN_ID_KLAYTN_BAOBAB) {
      return (amount: BigInt) => {
        dispatch(setIsApproving(true));
        return approveKlaytn(
          getTokenBridgeAddressForChain(chainId),
          tokenAddress,
          providerKaikas,
          BigNumber.from(amount),
          signerAddressKaikas
        ).then(
          () => {
            dispatch(setIsApproving(false));
            return Promise.resolve();
          },
          () => {
            dispatch(setIsApproving(false));
            return Promise.reject();
          }
        );
      };
    }
  },[
    chainId,
    tokenAddress,
    signer,
    dispatch,
    providerKaikas,
    signerAddressKaikas
  ]);

  return useMemo(
    () => ({
      sufficientAllowance,
      approveAmount,
      isAllowanceFetching,
      isApproveProcessing,
    }),
    [
      sufficientAllowance,
      approveAmount,
      isAllowanceFetching,
      isApproveProcessing,
    ]
  );
}
