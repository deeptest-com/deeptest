const fs = require('fs');
const axios = require('axios')
const FormData = require('form-data')

export async function uploadFile(url, token, path, params) {
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

    const resp = await axios.post(url, data, config)

    return resp.data
}
