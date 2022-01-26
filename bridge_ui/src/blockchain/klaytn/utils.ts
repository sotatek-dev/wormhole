import { arrayify, zeroPad } from "@ethersproject/bytes";

export function createNonce() {
    const nonceConst = Math.random() * 100000;
    const nonceBuffer = Buffer.alloc(4);
    nonceBuffer.writeUInt32LE(nonceConst, 0);
    return nonceBuffer;
  }

export const uint8ArrayToHex = (a: Uint8Array) =>
  Buffer.from(a).toString("hex");
export const nativeToHexStringKlaytn = (
  address: string | undefined,
  chain: any
) => {
  return uint8ArrayToHex(zeroPad(arrayify(address as any), 32));
}