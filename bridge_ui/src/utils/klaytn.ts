import {
  Implementation__factory
} from "@certusone/wormhole-sdk";
import klaytnBridgeImplementationAbi from "../blockchain/abi/BridgeImplementation.json";
import klaytnTokenImplementationAbi from "../blockchain/abi/TokenImplementation.json";
import klaytnNFTBridgeAbi from "../blockchain/abi/NFTBridge.json";
import klaytnERC721Abi from "../blockchain/abi/ERC721.json"
import { arrayify, zeroPad, formatUnits, BytesLike, Hexable } from 'ethers/lib/utils';
import caver from "../blockchain/klaytn/caver";
import { ChainId } from '@certusone/wormhole-sdk';
import { createNonce } from "../blockchain/klaytn/utils";
import { createNFTParsedTokenAccount, createParsedTokenAccount } from "../hooks/useGetSourceParsedTokenAccounts";

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
  wrappedAddress: string,
  lookupChainId: ChainId,
  provider?: any,
  tokenBridgeAddress?: string,
) {
  const isWrapped = await isWrappedAsset(wrappedAddress, provider, tokenBridgeAddress)
  if (isWrapped) {
      const contract = new provider.Contract(klaytnTokenImplementationAbi as any, wrappedAddress)
    const result = await contract.methods.nativeContract().call();
    const chainId = parseInt(await contract.methods.chainId().call());

      return {
          isWrapped: true,
          chainId: chainId as ChainId,
          assetAddress: arrayify(result),
      }
  }
  return {
      isWrapped: false,
      chainId: lookupChainId,
      assetAddress: zeroPad(arrayify(wrappedAddress as string), 32),
  }
}

export async function getOriginalAssetKlaytnNFT (
  tokenBridgeAddress: string,
  provider: any,
  wrappedAddress: string,
  tokenId: string,
  lookupChainId: ChainId,
  //tokenBridgeAddress, provider, wrappedAddress, tokenId, lookupChainId
) {
  const isWrapped = await isWrappedAsset(wrappedAddress, provider, tokenBridgeAddress)
  if (isWrapped) {
      const contract = new provider.Contract(klaytnTokenImplementationAbi as any, wrappedAddress)
      const result = await contract.methods.nativeContract().call();

      return {
          isWrapped: true,
          chainId: lookupChainId,
          assetAddress: arrayify(result),
          tokenId: tokenId
      }
  }
  return {
      isWrapped: false,
      chainId: lookupChainId,
      assetAddress: zeroPad(arrayify(wrappedAddress as string), 32),
      tokenId: tokenId
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
  const _recipientAddress = caver.utils.bytesToHex(recipientAddress as any)
  const _createNonce = caver.utils.bytesToHex(createNonce() as any)
  const result = await contract.methods
    .transferTokens(
      tokenAddress,
      amount,
      recipientChain,
      _recipientAddress,
      fee,
      _createNonce)
    .send({ from: signerAddress, gas: GAS_DEFAULT_KLAYTN })
  //address,uint256,uint16,bytes32,uint256,uint32
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
export async function klaytnTokenToParsedTokenAccount(
  tokenAddress: string,
  provider: any,
  signerAddress: string,
) {
  const contract = new provider.Contract(klaytnTokenImplementationAbi as any, tokenAddress)
  const decimals = await contract.methods.decimals().call()
  const balance = await contract.methods.balanceOf(signerAddress).call();
  const symbol = await contract.methods.symbol().call();
  const name = await contract.methods.name().call();

  return {
    address: contract?._address,
    decimals,
    balance: formatUnits(balance, decimals),
    symbol,
    name
  };
}

export async function getKlaytnNFT(
  tokenAddress: string,
  provider: any
) {
  const token = new provider.Contract(klaytnERC721Abi as any, tokenAddress)
  return token;
}

export async function isNFTKlaytn(contract: any) {
  const erc721 = "0x80ac58cd";
  const erc721metadata = "0x5b5e139f";
  const _erc721 = caver.utils.bytesToHex(arrayify(erc721) as any);
  const _erc721metadata = caver.utils.bytesToHex(arrayify(erc721metadata) as any);
  const supportsErc721 = await contract.methods
    .supportsInterface(_erc721)
    .call();
  const supportsErc721Metadata = await contract.methods
    .supportsInterface(_erc721metadata)
    .call();
  return supportsErc721 && supportsErc721Metadata;
}

export async function getKlaytnToken(
  tokenAddress: string,
  provider: any
) {
  const token = new provider.Contract(klaytnTokenImplementationAbi as any, tokenAddress)
  return token;
}

export async function klaytnNFTToNFTParsedTokenAccount(
  contract: any,
  tokenId: string,
  signerAddress: string
) {
  const decimals = 0;
  const ownerOf = await contract.methods.ownerOf(tokenId).call();
  const balance = ownerOf.toLowerCase() === signerAddress ? 1 : 0;
  const symbol = await contract.methods.symbol().call();
  const name = await contract.methods.name().call();
  const uri = await contract.methods.tokenURI(tokenId).call();
  return createNFTParsedTokenAccount(
    signerAddress,
    contract._address,
    balance.toString(),
    decimals,
    Number(formatUnits(balance, decimals)),
    formatUnits(balance, decimals),
    tokenId,
    symbol,
    name,
    uri
  );
}

export async function klaytnTokenToParsedTokenAccountNFT(
  token: any,
  signerAddress: string
) {
  const decimals = await token.methods.decimals().call();
  const balance = await token.methods.balanceOf(signerAddress).call();
  const symbol = await token.methods.symbol().call();
  const name = await token.methods.name().call();
  return createParsedTokenAccount(
    signerAddress,
    token?._address,
    balance.toString(),
    decimals,
    Number(formatUnits(balance, decimals)),
    formatUnits(balance, decimals),
    symbol,
    name
  );
} 

export async function redeemNftOnKlaytn (
  tokenBridgeAddress: string,
  provider: any,
  signerAddress: string,
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

export async function transferNFTFromKlaytn(
  tokenBridgeAddress: string,
  provider: any,
  signerAddressKaikas: string | undefined,
  tokenAddress: string,
  tokenId: string,
  recipientChain: any,
  recipientAddress: Uint8Array,
) {
  try {
    const contractERC721__factory = new provider.Contract(klaytnERC721Abi, tokenAddress);    
    const approved = await contractERC721__factory.methods
    .approve(tokenBridgeAddress, tokenId)
    .send({ from: signerAddressKaikas, gas: GAS_DEFAULT_KLAYTN });
    if (approved) {
      const contractNFTBridge__factory = new provider.Contract(klaytnNFTBridgeAbi, tokenBridgeAddress)
      const _recipientAddress = caver.utils.bytesToHex(recipientAddress as any)
      const _createNonce = caver.utils.bytesToHex(createNonce() as any)
      const result = await contractNFTBridge__factory.methods
        .transferNFT(tokenAddress, tokenId, recipientChain, _recipientAddress, _createNonce)
        .send({ from: signerAddressKaikas, gas: GAS_DEFAULT_KLAYTN })
      return result; 
    }
  } catch (error) {
    console.error(error)
  }
}

export function getEmitterAddressKlaytn(
  contractAddress: number | BytesLike | Hexable
) {
  return Buffer.from(zeroPad(arrayify(contractAddress), 32)).toString("hex");
}

// c96616e1 c6adb9658d8e6f589ac3b5a5490a90593e7e5accda2002e0d1d68f6b

// success
// 0xc96616e1
// 00000000000000000000000092556981a25918d141468a19462223d36cf1f70d 000000000000000000000000000000000000000000000000000000000000000a 0000000000000000000000000000000000000000000000000000000000001001 00000000000000000000000049114597ef077b8ddfa8c2be2dd35a1fe5c586c3 0000000000000000000000000000000000000000000000000000000000000064

// failed
// 0xc96616e1
// 000000000000000000000000ec990c8763cc90d2b72cf2806034da5002cf2e69 0000000000000000000000000000000000000000000000000000000000000001 0000000000000000000000000000000000000000000000000000000000000002 0000000000000000000000005c52ba41e7197136e679746b44c18885ada7b116 0000000000000000000000000000000000000000000000000000000086430100
