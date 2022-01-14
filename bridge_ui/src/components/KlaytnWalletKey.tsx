import { Typography } from "@material-ui/core";
import { useKaikasProvider } from "../contexts/KaikasProviderContext";
import ToggleConnectedButton from "./ToggleConnectedButton";

const KaikasWalletKey = () => {
  const { connect, disconnect, signerAddress, providerError } =
    useKaikasProvider();
  return (
    <>
      <ToggleConnectedButton
        connect={connect}
        disconnect={disconnect}
        connected={!!signerAddress}
        pk={signerAddress || ""}
      />
      {providerError ? (
        <Typography variant="body2" color="error">
          {providerError}
        </Typography>
      ) : null}
    </>
  );
};

export default KaikasWalletKey;
