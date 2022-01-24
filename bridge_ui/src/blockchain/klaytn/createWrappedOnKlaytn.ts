import {ChainId} from '@certusone/wormhole-sdk';
import Caver from 'caver-js';
import {arrayify, zeroPad} from 'ethers/lib/utils';
import klaytnBridgeImplementationAbi from '../abi/BridgeImplementation.json';
const KLAYTN_PROVIDER_API = process.env.REACT_APP_KLAYTN_PROVIDER_API;
const KEY_KLAYTN = process.env.REACT_APP_KEY_KLAYTN;
const GAS_PRICE_KLAYTN = "25000000000";

export async function createWrappedOnKlaytn (
   tokenBridgeAddress: string,
   signerAddress: any,
   signedVAA: Uint8Array
) {
    await window.klaytn.enable();
    const caver =  new Caver(window.klaytn);
    const contract = new caver.klay.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
    const encodeVM = caver.utils.bytesToHex(signedVAA as any)
    const result = await contract.methods.updateWrapped(
      encodeVM
    ).send({
        from: signerAddress,
        gas: 3000000
    });
    console.log('result: ', result);
    return result;
}