import klaytnBridgeImplementationAbi from "../blockchain/abi/BridgeImplementation.json";
import klaytnTokenImplementationAbi from "../blockchain/abi/TokenImplementation.json";
import Caver from 'caver-js';
import {arrayify, zeroPad} from 'ethers/lib/utils';

import { ChainId } from '@certusone/wormhole-sdk';
import { createNonce } from "../blockchain/klaytn/utils";

const caver = new Caver(window.klaytn);
const TOKEN_BRIDGE_ADDRESS_KLAYTN = process.env.REACT_APP_TOKEN_BRIDGE;
const KLAYTN_PROVIDER_API = process.env.REACT_APP_KLAYTN_PROVIDER_API;
const KEY_KLAYTN = process.env.REACT_APP_KEY_KLAYTN;

const GAS_PRICE_KLAYTN = "25000000000";

export async function getForeignAssetKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  signerAddressKaikas: string | undefined,
  originChain: ChainId,
  originAsset: Uint8Array
) {
  try {
    console.log("xx");
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
  .send({from: signerAddress, gas: 3000000 })
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
    gas: 3000000
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
    gas: 3000000
  });
  return result;
}

export default async function isWrappedAsset(
    address?: string, 
) {
    
    const caver = new Caver(KLAYTN_PROVIDER_API);

    const contract = new caver.klay.Contract(klaytnBridgeImplementationAbi as any, TOKEN_BRIDGE_ADDRESS_KLAYTN, {
        from: KEY_KLAYTN,
        gasPrice: GAS_PRICE_KLAYTN
    })

    const result = await contract.methods.isWrappedAsset(address).call();
    
    return result;
}

export async function getOriginalAssetKlaytn (
  wrappedAddress?: string,
  lookupChainId?: ChainId
) {
  const isWrapped = await isWrappedAsset(wrappedAddress)
  if (isWrapped) {
      const caver = new Caver(KLAYTN_PROVIDER_API);
      const contract = new caver.klay.Contract(klaytnTokenImplementationAbi as any, wrappedAddress, {
          from: KEY_KLAYTN,
          gasPrice: GAS_PRICE_KLAYTN
      })
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
    gas: 3000000
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
    gas: 3000000
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