const express = require('express');
const app = express();
const path = require('path');
const PORT = process.env.REACT_APP_PORT_SERVER || 3001;

app.use(express.static(path.join(__dirname, 'build')));

app.listen(PORT, () => {
    console.log(`App listening on port ${PORT}`);
});
