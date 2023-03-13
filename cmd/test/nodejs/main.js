import got from 'got';
import crypto from 'crypto';
import fs from 'fs';
import os from "os";

// const t = await got.get('https://dl.cnezsoft.com/zd/2.3/darwin/zd.zip.md5').text();
// console.log(t);
//
// const buffer = fs.readFileSync('/Users/aaron/zd/tmp/download/2.3.zip');
// const hash = crypto.createHash('md5');
// hash.update(buffer, 'utf8');
// const md5 = hash.digest('hex');
// console.log(md5);

// fse.copySync('/Users/aaron/zd/tmp/download/extracted/runtime',
//     '/Users/aaron/rd/project/gudi/deeptest/client/out/darwin/deeptest.app/Contents/Resources/bin/runtime')

// console.log(parseFloat('1.1.1'))
// export const DEBUG = true //process.env.NODE_ENV === 'development';
// console.log(DEBUG)

// export function getPlatform() {
//     let platform = os.platform(); // 'darwin', 'linux', 'win32'
//
//     if (platform === 'win32' && ['arm64', 'ppc64', 'x64', 's390x'].includes(os.arch())) {
//         platform = 'win64'
//     }
//
//     return platform
// }
//
// console.log(getPlatform())

console.log('server.exe  4568'.split(/\s/).length)