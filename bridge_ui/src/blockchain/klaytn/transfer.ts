import klaytnBridgeImplementationAbi from '../abi/BridgeImplementation.json';
import { createNonce } from './utils';
import Caver from 'caver-js';

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
  .send({ from: signerAddress, gas: 3000000 })
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
  const caver =  new Caver(window.klaytn);
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress)
  const fee = 0;
  const encodeVM = caver.utils.bytesToHex(recipientAddress as any)
  const result = await contract.methods.wrapAndTransferETH(recipientChain, encodeVM, fee, createNonce())
    .send({ from: signerAddress, value: amount, gas: 3000000 })
  return result;
}
