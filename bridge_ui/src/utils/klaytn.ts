import klaytnBridgeImplementationAbi from "../blockchain/abi/BridgeImplementation.json";
import Caver from 'caver-js';

import { ChainId } from '@certusone/wormhole-sdk';

const caver = new Caver(window.klaytn);

export async function getForeignAssetKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  signerAddressKaikas: string | undefined,
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