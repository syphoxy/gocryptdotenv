# gocryptdotenv

This library is a wrapper for godotenv that facilitates encrypting the values in an `.env` file.

## Build

Building is easy!

    $ go build cmd/crypt-env.go

Now you can use `./crypt-env`!

## Encrypt

There are three flags. `-f FILE` to specify the path to the plaintext .env file. `-t encrypt` to specify that you want to encrypt. `-k KEY` to specify the key that you want to use to encrypt that file with.

    $ cat fixtures/plaintext.env
    GO_ENVIRONMENT='production'
    GO_API_KEY='a_really_long_api_key'
    GO_API_SECRET='another_really_long_secret_key'
    GO_FOO='bar'
    GO_HELLO='world'

    $ ./crypt-env -f fixtures/plaintext.env -t encrypt -k testkey
    GO_HELLO='d57f434664c947711a700f9e55c4592293fc89d8fa6c8863ee09282c3a6a4587 70a95fd5db2da4d268d0ba209764d6de40ef6a9bcc440e8b97777f0367d68a5532b63dc51b63e6b2b02bc54290'
    GO_ENVIRONMENT='5e4747996642de699fed2d3f898eec36240b1890dbb569f27b2e6d933284b8b3 119e05a888f77ac899c8d7d2a3decfb274eb9d8ade4349397ae09ca2aecf7857076150a64d17a28068e8edc4db9b5e485957'
    GO_API_KEY='6b3a60c24364360a2685c5126135ca2c413ba407a33f7e77750a5ce07c2b16c1 b2c23a9be71ac816461961b379f65c3c6ab430dab7d20dcd52e35c108e3e68898fb91e7aecc8be30ea4b052713983e1d6a484daf47823411fa816b356e'
    GO_API_SECRET='d3729c3960aa5cf7e18883641a067bd48df44c08abf66424844a024a924d9cae bc2b8af3cc6aae282b12a5c29a375775068cf82ba0ebf9879edfbb9d1bbe85a22a3922b7da173f963accc032e0b4d1e0a9394ddd917d60d75459a27b9117d726d861c8637ff6'
    GO_FOO='7c09b2f9cc99444bab843417a941b1d90d33b3bb61fa199e4f3a927e89f790cd aaf5828ffe3b5e5c8745620c3d97fd5772c2be1df5111d951b6ade7be20291c607ae619e9578468f49afdb'

## Decrypt

Same arguments as the encrypt operation except `-t decrypt`.

    $ cat fixtures/ciphertext.env
    GO_HELLO='fc2d579c11118f591d6fbbc033dfabbf9e0f4e218e30cceaf8fadd05136781c5 8767534af80ed5265c1df0288bc68e7bec07e44398782bdce68e5b327af4572def82b7014d11e4ce8f63278dc0'
    GO_ENVIRONMENT='d13937c20725c93861211bc2a92b166e868db3a2f8e60bd83f2164266060a708 e2be8662f93eef1a2ca0d4976ba31e794421ff4bd3661a02554b71ea64ddfd6aedfb010583a4db8c487d634369f037d0e32b'
    GO_API_KEY='6d07989ae1fdf161a0a44ffffd563aacc00e12c0ac2512e687266e2a0543440d 332307bbf9df16bdbf6ec38ddd73b8eae6aa115a087126129261fed9964b731f61659b85b9ac64c7606bb7a89813509f7bf5bd244552d95000c1ed0f6f'
    GO_API_SECRET='86bc2b6eec1be760c080af620f7a1c0b38efa6313585d63cf1540a6b7db767c8 faf55df0b68e1c06b81275c4d1312c5d44c8550cfc6f21d1550e7bb661efa55dfee9d409fd40975f79d88e44fc7e35267b055b6a39e979d5de680edf4994fd45e1a401bb0755'
    GO_FOO='abb1dfd63c033857c2682a1b76e02cfd2967f8ea1e761b9b1ea42cf68b098856 1387b10b07f1c7fead6b4422a539878612157fc9b29d34eaf91ea5eaae263d4d6891d3e76e576f295f41c8'

    $ ./crypt-env -f fixtures/ciphertext.env -t decrypt -k testkey
    GO_FOO='bar'
    GO_HELLO='world'
    GO_ENVIRONMENT='production'
    GO_API_KEY='a_really_long_api_key'
    GO_API_SECRET='another_really_long_secret_key'
