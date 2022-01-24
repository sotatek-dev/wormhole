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
    const data = caver.klay.abi.encodeFunctionCall(
        {
            inputs: [
              {
                internalType: "bytes",
                name: "encodedVm",
                type: "bytes"
              }
            ],
            name: "createWrapped",
            outputs: [
              {
                internalType: "address",
                name: "token",
                type: "address"
              }
            ],
            stateMutability: "nonpayable",
            type: "function"
          },
        [signedVAA]
      )
     const result = caver.klay.sendTransaction({
        type: 'SMART_CONTRACT_EXECUTION',
        from: signerAddress,
        to: tokenBridgeAddress,
        gas: '8000000',
        data
    })
    // const contract = new caver.klay.Contract(klaytnBridgeImplementationAbi as any, tokenBridgeAddress);
    // const result = await contract.methods.createWrapped(
    //     signedVAA
    // ).send({
    //     from: signerAddress,
    // });
    console.log('result: ', result);
    return result;
}