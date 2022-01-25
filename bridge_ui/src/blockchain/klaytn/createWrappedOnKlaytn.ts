import caver from './caver';
import klaytnBridgeImplementationAbi from '../abi/BridgeImplementation.json';

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
    gas: 3000000
  });
  return result;
}