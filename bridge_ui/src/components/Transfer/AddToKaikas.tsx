import { CHAIN_ID_KLAYTN_BAOBAB } from "@certusone/wormhole-sdk";
import { Button, makeStyles } from "@material-ui/core";
import { useCallback } from "react";
import { useSelector } from "react-redux";
import { useKaikasProvider } from "../../contexts/KaikasProviderContext";
import {
  selectTransferSourceParsedTokenAccount,
  selectTransferTargetAsset,
  selectTransferTargetChain,
} from "../../store/selectors";
import { getEvmChainId } from "../../utils/consts";
import { klaytnTokenToParsedTokenAccount } from "../../utils/klaytn";

const useStyles = makeStyles((theme) => ({
  addButton: {
    display: "block",
    margin: `${theme.spacing(1)}px auto 0px`,
  },
}));

export default function AddToKaikas() {
  const classes = useStyles();
  const sourceParsedTokenAccount = useSelector(
    selectTransferSourceParsedTokenAccount
  );
  const targetChain = useSelector(selectTransferTargetChain);
  const targetAsset = useSelector(selectTransferTargetAsset);
  
  const {provider: providerKaikas, signerAddress: signerAddressKaikas, chainId: chainIdKaikas} = useKaikasProvider();
  const hasCorrectEvmNetwork = chainIdKaikas === getEvmChainId(targetChain);
  const handleClick = useCallback(() => {
    if (providerKaikas && targetAsset && signerAddressKaikas && hasCorrectEvmNetwork) {
      (async () => {
        try {
            const result = await klaytnTokenToParsedTokenAccount(targetAsset, providerKaikas, signerAddressKaikas);
            const {decimals, symbol, address} = result;
            const {klaytn} = window;
            klaytn.sendAsync(
                {
                    method: 'wallet_watchAsset',
                    params: {
                      type: 'ERC20', 
                      options: {
                        address, 
                        symbol: symbol ||  sourceParsedTokenAccount?.symbol, 
                        decimals,
                        image: "" 
                      }
                    },
                    id: Math.round(Math.random() * 100000)
                  },
                  (err: any, result: any) => console.log(err, result)
            )
        } catch (e) {
          console.error(e);
        }
      })();
    }
  }, [
    targetAsset,
    hasCorrectEvmNetwork,
    providerKaikas,
    signerAddressKaikas,
    sourceParsedTokenAccount
  ]);
  return providerKaikas &&
    signerAddressKaikas &&
    targetAsset &&
    targetChain === CHAIN_ID_KLAYTN_BAOBAB &&
    hasCorrectEvmNetwork ? (
    <Button
      onClick={handleClick}
      size="small"
      variant="outlined"
      className={classes.addButton}
    >
      Add to Kaikas
    </Button>
  ) : null;
}
