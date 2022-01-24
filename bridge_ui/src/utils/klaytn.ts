import klaytnBridgeImplementationAbi from "../blockchain/abi/BridgeImplementation.json";

import { ChainId } from '@certusone/wormhole-sdk';

export async function getForeignAssetKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  originChain: ChainId,
  originAsset: Uint8Array
) {
  try {
    const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
    const result = await contract.methods.wrappedAsset(originChain, originAsset).call();
    console.log('result: ', result);
    return result;

  } catch (error) {
    return null;
  }
}