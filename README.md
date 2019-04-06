# Crypta

Crypta is another data encryption tool that uses RSA-AES encryption to securely transfer your data.

## Getting Started

### Installation

#### Linux

#### MacOS

### Usage

1. Generate a key pair using `openssl`
   
    ```bash
    openssl genrsa -out private.key 2048
    openssl rsa -pubout -in private.key -out public.key
    
    # Testing your key pair with crypta. This should print 'Hello World' as the output
    echo "Hello World" | crypta -key public.key | crypta -d -key private.key
    ```
   If you don't want to use `-key` flag you can use `CRYPTA_PRIVATE_KEY` and `CRYPTA_PUBLIC_KEY` environment variables to set the path for keys. 
   However priority is given to the `-key` flag if you specify both. 
   
2. Encrypt using public key

    ```bash
    # Encrypt message.txt using public key and save it as payload.bin
    crypta -in message.txt -key public.key -out payload.bin
 
    # Encrypt message.txt using public key and save it as base64 encoded payload.txt
    crypta -in message.txt -key public.key -out payload.txt -base64
 
    # Read data from stdin pipe and encrypt
    echo "Hello World" | crypta -key public.key -out payload.bin

    # Read data from stdin and encrypt (press Ctrl+D after entering a secret)
    crypta -key public.key -out payload.bin
    ```

3. Decrypt using private key

    ```bash
    # Decrypt payload.bin using private key and save it as message.txt
    crypta -in payload.bin -key private.key -out message.txt -d
 
    # Decrypt base64 encoded payload.txt using private key and save it as message.txt
    crypta -in payload.txt -key private.key -out message.txt -d -base64
 
    # Read data from stdin pipe and decrypt
    cat payload.bin | crypta -key private.key -out message.txt -d

    # Output decrypt data to stdout
    crypta -in payload.bin -key private.key -d
    ```
    
4. Run `crypta --help` for more information

    ```bash
    Usage of crypta:
      -base64
            Use base64 encoding to write output (encrypt mode), and base64 decoding to read input (decrypt mode)
      -d    Decrypt mode
      -in string
            Input file path. Uses standard input if not provided
      -key string
            Public key file path for encrypting or private key file path for decrypting
      -out string
            Output file path. Uses standard output if not provided
      -version
            output version information
    
    ```
