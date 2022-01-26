import {
  Implementation__factory
} from "@certusone/wormhole-sdk";
import klaytnBridgeImplementationAbi from "../blockchain/abi/BridgeImplementation.json";
import klaytnTokenImplementationAbi from "../blockchain/abi/TokenImplementation.json";
import {arrayify, zeroPad} from 'ethers/lib/utils';
import caver from "../blockchain/klaytn/caver";
import { ChainId } from '@certusone/wormhole-sdk';
import { createNonce } from "../blockchain/klaytn/utils";

export const GAS_DEFAULT_KLAYTN = 3000000;

export function parseSequenceFromLogKlaytn (
  receipt: any,
  bridgeAddress: string
): string {
  let bridge = {} as any;
  const keyEvents = Object.keys(receipt?.events);
  for(let i = 0; i <keyEvents.length; i++) {
    const bridgeItem = receipt?.events[keyEvents[i]];
    if (bridgeItem?.address === bridgeAddress) {
      bridge = bridgeItem
    }
  }
  
  const bridgeLog = {
    ...bridge,
    data: bridge?.raw?.data,
    topics: bridge?.raw?.topics
  }
  const {
    args: { sequence },
  } = Implementation__factory.createInterface().parseLog(bridgeLog);
  return sequence.toString();
}

export async function getForeignAssetKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  originChain: ChainId,
  originAsset: Uint8Array
) {
  try {
    const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
    const _originAsset = caver.utils.bytesToHex(originAsset as any)
    const result = await contract.methods.wrappedAsset(originChain, _originAsset).call();
    return result;

  } catch (error: any) {
    throw new Error(error.message)
  }
}

export async function attestFromKlaytn (
  tokenBridgeAddress: string,
  provider: any,
  tokenAddress: string,
  signerAddress: any,
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress)
  console.log(tokenAddress);
  
  const result = await contract.methods.attestToken(tokenAddress, createNonce())
  .send({from: signerAddress, gas: GAS_DEFAULT_KLAYTN })
  return result;
  
}

export async function createWrappedOnKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  signerAddress: string | undefined,
  signedVAA: Uint8Array
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
  const encodeVM = caver.utils.bytesToHex(signedVAA as any)
  const result = await contract.methods.createWrapped(
    encodeVM
  ).send({
    from: signerAddress,
    gas: GAS_DEFAULT_KLAYTN
  });
  return result;
}

export async function updateWrappedOnKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  signerAddress: string | undefined,
  signedVAA: Uint8Array
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
  const encodeVM = caver.utils.bytesToHex(signedVAA as any)
  const result = await contract.methods.updateWrapped(
    encodeVM
  ).send({
    from: signerAddress,
    gas: GAS_DEFAULT_KLAYTN
  });
  return result;
}

export default async function isWrappedAsset(
    address?: string, 
    provider?: any,
    tokenBridgeAddress?: string
) {
    
    const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress)

    const result = await contract.methods.isWrappedAsset(address).call();
    
    return result;
}

export async function getOriginalAssetKlaytn (
  wrappedAddress?: string,
  lookupChainId?: ChainId,
  provider?: any,
  tokenBridgeAddress?: string,
) {
  const isWrapped = await isWrappedAsset(wrappedAddress, provider, tokenBridgeAddress)
  if (isWrapped) {
      const contract = new provider.Contract(klaytnTokenImplementationAbi as any, wrappedAddress)
      const result = await contract.methods.nativeContract().call();

      return {
          isWrapped: true,
          chainId: lookupChainId,
          assetAddress: arrayify(result),
      }
  }
  return {
      isWrapped: false,
      chainId: lookupChainId,
      assetAddress: zeroPad(arrayify(wrappedAddress as string), 32),
  }
}

export async function redeemOnKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  signerAddress: string | undefined,
  signedVAA: Uint8Array
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
  const encodeVM = caver.utils.bytesToHex(signedVAA as any)
  const result = await contract.methods.completeTransfer(
    encodeVM
  ).send({
    from: signerAddress,
    gas: GAS_DEFAULT_KLAYTN
  });
  return result;
}

export async function redeemOnKlaytnNative(
  tokenBridgeAddress: string,
  provider: any,
  signerAddress: string | undefined,
  signedVAA: Uint8Array
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
  const encodeVM = caver.utils.bytesToHex(signedVAA as any)
  const result = await contract.methods.completeTransferAndUnwrap(
    encodeVM
  ).send({
    from: signerAddress,
    gas: GAS_DEFAULT_KLAYTN
  });
  return result;
}


export async function transferFromKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  tokenAddress: string,
  amount: any,
  recipientChain: any,
  recipientAddress: Uint8Array,
  signerAddress: any,
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress)
  const fee = 0;
  const result = await contract.methods
  .transferTokens(tokenAddress, amount, recipientChain, recipientAddress, fee, createNonce())
  .send({ from: signerAddress, gas: GAS_DEFAULT_KLAYTN })
  return result;
}

export async function transferFromKlaytnNative(
  tokenBridgeAddress: string,
  provider: any,
  amount: any,
  recipientChain: any,
  recipientAddress: Uint8Array,
  signerAddress: any,
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress)
  const fee = 0;
  const encodeVM = caver.utils.bytesToHex(recipientAddress as any)
  const result = await contract.methods.wrapAndTransferETH(recipientChain, encodeVM, fee, createNonce())
    .send({ from: signerAddress, value: amount, gas: GAS_DEFAULT_KLAYTN })
  return result;
}