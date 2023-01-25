const fs = require('fs');
const axios = require('axios')

export async function uploadDatapoolFile(url, token, path, params) {
    const data = new FormData();

    data.append('file', fs.createReadStream(path));
    // data.append('file', fs.readFileSync(path));

    Object.keys(params).forEach(key=>{
        data.append(key, params[key]);
    })

    const config = {
        headers: {
            'Content-Type': 'multipart/form-data',
            'Authorization': 'Bearer ' + token
        },
        onUploadProgress: function(progressEvent) {
            const percentCompleted = Math.round( (progressEvent.loaded * 100) / progressEvent.total );
        }
    };

    const res = await axios.post(url, data, config)
    console.log('===', res)
}
