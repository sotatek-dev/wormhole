import caver from './caver';
import klaytnBridgeImplementationAbi from '../abi/BridgeImplementation.json';

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
    gas: 3000000
  });
  return result;
}