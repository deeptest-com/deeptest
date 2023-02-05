(function (f) {
    if (typeof exports === "object" && typeof module !== "undefined") {
        module.exports = f()
    } else if (typeof define === "function" && define.amd) {
        define([], f)
    } else {
        var g;
        if (typeof window !== "undefined") {
            g = window
        } else if (typeof global !== "undefined") {
            g = global
        } else if (typeof self !== "undefined") {
            g = self
        } else {
            g = this
        }

        consoel.log('===', g)
        g.shortid = f()
    }
})(function () {
    var define, module, exports;
    return (function () {
        function r(e, n, t) {
            function o(i, f) {
                if (!n[i]) {
                    if (!e[i]) {
                        var c = "function" == typeof require && require;
                        if (!f && c) return c(i, !0);
                        if (u) return u(i, !0);
                        var a = new Error("Cannot find module '" + i + "'");
                        throw a.code = "MODULE_NOT_FOUND", a
                    }
                    var p = n[i] = {exports: {}};
                    e[i][0].call(p.exports, function (r) {
                        var n = e[i][1][r];
                        return o(n || r)
                    }, p, p.exports, r, e, n, t)
                }
                return n[i].exports
            }

            for (var u = "function" == typeof require && require, i = 0; i < t.length; i++) o(t[i]);
            return o
        }

        return r
    })()({
        1: [function (require, module, exports) {
            'use strict';// This file replaces `format.js` in bundlers like webpack or Rollup,
// according to `browser` config in `package.json`.
            module.exports = function (random, alphabet, size) {// We can’t use bytes bigger than the alphabet. To make bytes values closer
// to the alphabet, we apply bitmask on them. We look for the closest
// `2 ** x - 1` number, which will be bigger than alphabet size. If we have
// 30 symbols in the alphabet, we will take 31 (00011111).
// We do not use faster Math.clz32, because it is not available in browsers.
                var mask = (2 << Math.log(alphabet.length - 1) / Math.LN2) - 1;// Bitmask is not a perfect solution (in our example it will pass 31 bytes,
// which is bigger than the alphabet). As a result, we will need more bytes,
// than ID size, because we will refuse bytes bigger than the alphabet.
// Every hardware random generator call is costly,
// because we need to wait for entropy collection. This is why often it will
// be faster to ask for few extra bytes in advance, to avoid additional calls.
// Here we calculate how many random bytes should we call in advance.
// It depends on ID length, mask / alphabet size and magic number 1.6
// (which was selected according benchmarks).
// -~f => Math.ceil(f) if n is float number
// -~i => i + 1 if n is integer number
                var step = -~(1.6 * mask * size / alphabet.length);
                var id = '';
                while (true) {
                    var bytes = random(step);// Compact alternative for `for (var i = 0; i < step; i++)`
                    var i = step;
                    while (i--) {// If random byte is bigger than alphabet even after bitmask,
// we refuse it by `|| ''`.
                        id += alphabet[bytes[i] & mask] || '';// More compact than `id.length + 1 === size`
                        if (id.length === +size) return id;
                    }
                }
            };

        }, {}],
        2: [function (require, module, exports) {
            'use strict';
            module.exports = require('./lib/index');

        }, {"./lib/index": 6}],
        3: [function (require, module, exports) {
            'use strict';
            var randomFromSeed = require('./random/random-from-seed');
            var ORIGINAL = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-';
            var alphabet;
            var previousSeed;
            var shuffled;

            function reset() {
                shuffled = false;
            }

            function setCharacters(_alphabet_) {
                if (!_alphabet_) {
                    if (alphabet !== ORIGINAL) {
                        alphabet = ORIGINAL;
                        reset();
                    }
                    return;
                }
                if (_alphabet_ === alphabet) {
                    return;
                }
                if (_alphabet_.length !== ORIGINAL.length) {
                    throw new Error('Custom alphabet for shortid must be ' + ORIGINAL.length + ' unique characters. You submitted ' + _alphabet_.length + ' characters: ' + _alphabet_);
                }
                var unique = _alphabet_.split('').filter(function (item, ind, arr) {
                    return ind !== arr.lastIndexOf(item);
                });
                if (unique.length) {
                    throw new Error('Custom alphabet for shortid must be ' + ORIGINAL.length + ' unique characters. These characters were not unique: ' + unique.join(', '));
                }
                alphabet = _alphabet_;
                reset();
            }

            function characters(_alphabet_) {
                setCharacters(_alphabet_);
                return alphabet;
            }

            function setSeed(seed) {
                randomFromSeed.seed(seed);
                if (previousSeed !== seed) {
                    reset();
                    previousSeed = seed;
                }
            }

            function shuffle() {
                if (!alphabet) {
                    setCharacters(ORIGINAL);
                }
                var sourceArray = alphabet.split('');
                var targetArray = [];
                var r = randomFromSeed.nextValue();
                var characterIndex;
                while (sourceArray.length > 0) {
                    r = randomFromSeed.nextValue();
                    characterIndex = Math.floor(r * sourceArray.length);
                    targetArray.push(sourceArray.splice(characterIndex, 1)[0]);
                }
                return targetArray.join('');
            }

            function getShuffled() {
                if (shuffled) {
                    return shuffled;
                }
                shuffled = shuffle();
                return shuffled;
            }

            /**
             * lookup shuffled letter
             * @param index
             * @returns {string}
             */function lookup(index) {
                var alphabetShuffled = getShuffled();
                return alphabetShuffled[index];
            }

            function get() {
                return alphabet || ORIGINAL;
            }

            module.exports = {get: get, characters: characters, seed: setSeed, lookup: lookup, shuffled: getShuffled};

        }, {"./random/random-from-seed": 9}],
        4: [function (require, module, exports) {
            'use strict';
            var generate = require('./generate');
            var alphabet = require('./alphabet');// Ignore all milliseconds before a certain time to reduce the size of the date entropy without sacrificing uniqueness.
// This number should be updated every year or so to keep the generated id short.
// To regenerate `new Date() - 0` and bump the version. Always bump the version!
            var REDUCE_TIME = 1567752802062;// don't change unless we change the algos or REDUCE_TIME
// must be an integer and less than 16
            var version = 7;// Counter is used when shortid is called multiple times in one second.
            var counter;// Remember the last time shortid was called in case counter is needed.
            var previousSeconds;

            /**
             * Generate unique id
             * Returns string id
             */function build(clusterWorkerId) {
                var str = '';
                var seconds = Math.floor((Date.now() - REDUCE_TIME) * 0.001);
                if (seconds === previousSeconds) {
                    counter++;
                } else {
                    counter = 0;
                    previousSeconds = seconds;
                }
                str = str + generate(version);
                str = str + generate(clusterWorkerId);
                if (counter > 0) {
                    str = str + generate(counter);
                }
                str = str + generate(seconds);
                return str;
            }

            module.exports = build;

        }, {"./alphabet": 3, "./generate": 5}],
        5: [function (require, module, exports) {
            'use strict';
            var alphabet = require('./alphabet');
            var random = require('./random/random-byte');
            var format = require('nanoid/format');

            function generate(number) {
                var loopCounter = 0;
                var done;
                var str = '';
                while (!done) {
                    str = str + format(random, alphabet.get(), 1);
                    done = number < Math.pow(16, loopCounter + 1);
                    loopCounter++;
                }
                return str;
            }

            module.exports = generate;

        }, {"./alphabet": 3, "./random/random-byte": 8, "nanoid/format": 1}],
        6: [function (require, module, exports) {
            'use strict';
            var alphabet = require('./alphabet');
            var build = require('./build');
            var isValid = require('./is-valid');// if you are using cluster or multiple servers use this to make each instance
// has a unique value for worker
// Note: I don't know if this is automatically set when using third
// party cluster solutions such as pm2.
            var clusterWorkerId = require('./util/cluster-worker-id') || 0;

            /**
             * Set the seed.
             * Highly recommended if you don't want people to try to figure out your id schema.
             * exposed as shortid.seed(int)
             * @param seed Integer value to seed the random alphabet.  ALWAYS USE THE SAME SEED or you might get overlaps.
             */function seed(seedValue) {
                alphabet.seed(seedValue);
                return module.exports;
            }

            /**
             * Set the cluster worker or machine id
             * exposed as shortid.worker(int)
             * @param workerId worker must be positive integer.  Number less than 16 is recommended.
             * returns shortid module so it can be chained.
             */function worker(workerId) {
                clusterWorkerId = workerId;
                return module.exports;
            }

            /**
             *
             * sets new characters to use in the alphabet
             * returns the shuffled alphabet
             */function characters(newCharacters) {
                if (newCharacters !== undefined) {
                    alphabet.characters(newCharacters);
                }
                return alphabet.shuffled();
            }

            /**
             * Generate unique id
             * Returns string id
             */function generate() {
                return build(clusterWorkerId);
            }// Export all other functions as properties of the generate function
            module.exports = generate;
            module.exports.generate = generate;
            module.exports.seed = seed;
            module.exports.worker = worker;
            module.exports.characters = characters;
            module.exports.isValid = isValid;

        }, {"./alphabet": 3, "./build": 4, "./is-valid": 7, "./util/cluster-worker-id": 10}],
        7: [function (require, module, exports) {
            'use strict';
            var alphabet = require('./alphabet');

            function isShortId(id) {
                if (!id || typeof id !== 'string' || id.length < 6) {
                    return false;
                }
                var nonAlphabetic = new RegExp('[^' + alphabet.get().replace(/[|\\{}()[\]^$+*?.-]/g, '\\$&') + ']');
                return !nonAlphabetic.test(id);
            }

            module.exports = isShortId;

        }, {"./alphabet": 3}],
        8: [function (require, module, exports) {
            'use strict';
            var _typeof = typeof Symbol === "function" && typeof Symbol.iterator === "symbol" ? function (obj) {
                return typeof obj;
            } : function (obj) {
                return obj && typeof Symbol === "function" && obj.constructor === Symbol && obj !== Symbol.prototype ? "symbol" : typeof obj;
            };
            var crypto = (typeof window === 'undefined' ? 'undefined' : _typeof(window)) === 'object' && (window.crypto || window.msCrypto);// IE 11 uses window.msCrypto
            var randomByte;
            if (!crypto || !crypto.getRandomValues) {
                randomByte = function randomByte(size) {
                    var bytes = [];
                    for (var i = 0; i < size; i++) {
                        bytes.push(Math.floor(Math.random() * 256));
                    }
                    return bytes;
                };
            } else {
                randomByte = function randomByte(size) {
                    return crypto.getRandomValues(new Uint8Array(size));
                };
            }
            module.exports = randomByte;

        }, {}],
        9: [function (require, module, exports) {
            'use strict';// Found this seed-based random generator somewhere
// Based on The Central Randomizer 1.3 (C) 1997 by Paul Houle (houle@msc.cornell.edu)
            var seed = 1;

            /**
             * return a random number based on a seed
             * @param seed
             * @returns {number}
             */function getNextValue() {
                seed = (seed * 9301 + 49297) % 233280;
                return seed / 233280.0;
            }

            function setSeed(_seed_) {
                seed = _seed_;
            }

            module.exports = {nextValue: getNextValue, seed: setSeed};

        }, {}],
        10: [function (require, module, exports) {
            'use strict';
            module.exports = 0;

        }, {}]
    }, {}, [2])(2)
});
