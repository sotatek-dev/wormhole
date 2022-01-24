import klaytnBridgeImplementationAbi from '../abi/BridgeImplementation.json';
import { createNonce } from './utils';

export async function attestFromKlaytn (
  tokenBridgeAddress: string,
  provider: any,
  tokenAddress: string,
  signerAddress: any,
) {
  const contract = new provider.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress)
  console.log(tokenAddress);
  
  const result = await contract.methods.attestToken(tokenAddress, createNonce())
  .send({from: signerAddress, gas: 3000000 })
  return result;
  
}
