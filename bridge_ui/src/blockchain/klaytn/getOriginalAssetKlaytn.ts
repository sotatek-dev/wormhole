import {ChainId} from '@certusone/wormhole-sdk';
import Caver from 'caver-js';
import {arrayify, zeroPad} from 'ethers/lib/utils';
import klaytnTokenImplementationAbi from '../abi/TokenImplementation.json';
import isWrappedAsset from "./isWrappedAsset";

const KLAYTN_PROVIDER_API = process.env.REACT_APP_KLAYTN_PROVIDER_API;
const KEY_KLAYTN = process.env.REACT_APP_KEY_KLAYTN;
const GAS_PRICE_KLAYTN = "25000000000";
export async function getOriginalAssetKlaytn (
    wrappedAddress?: string,
    lookupChainId?: ChainId
) {
    console.log(wrappedAddress);
    
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