---                                                                                   
title: Handin 3 - Disys
author: Thea Hende 202105228
        Helena Cooper 201906086,
        Timmi Andersen 202105859
date: 29/09/2023
--- 


## Exercise 6.1 (Enc-then-sign or Sign-then-encrypt)

### Opgave beskrivelse til 6.1

You are consultant to a company that offers access to a database $D$. Different users have access to different parts of the database and moreover, the database contains sensitive information. Therefore the security policy says:

- When a request for information arrives, $D$ should be able to determine which user sent the request.
- A user should not be able to get get information about which data other users asked for.

We can assume that every user $A$ has a private RSA key $s k_A$ of his own, and the database stores a list of public keys of all users. Also all users know the public key $p k_D$ of $D$.

The threat model assumes that the database is sufficiently protected that it will not be hacked, but network traffic can attacked and manipulated.

One suggestion to ensure that the policy is followed is: when user $A$ wants to send a request $R$ to $D$, he will send

$$
E_{p k_D}(R), S_{s k_A}\left(E_{p k_D}(R)\right), A
$$

to $D$, that is, encrypt the request under $D$ 's public key, append his signature on the encrypted request, and then append his username. $D$ will look up $p k_A$, verify the signature, and if it is correct, will decrypt the request, and return an answer of form $E_{p k_A}$ (data), that is, the answer to the request, encrypted under $A$ 's public key.

Another suggestion goes as follows: when user $A$ wants to send a request $R$ to $D$, he will send

$$
E_{p k_D}\left(R, S_{s k_A}(R)\right), A
$$

to $D$, that is, append his signature on the request, encrypt this under $D$ 's public key, and then append his username. $D$ will decrypt the request and signature, look up $p k_A$, verify the signature, and if it is correct, will return an answer of form $E_{p k_A}$ (data), that is, the answer to the request, encrypted under $A$ 's public key.

One of these solutions fails to satisfy the security policy. Which one, and why? are there any general conclusions one could draw from this example?

### Besvarelse til 6.1

## Exercise 5.11

### Opgave beskrivelse til 5.11 (RSA and AES encryption)

This exercise has two parts. The first asks you to implement RSA yourself. The second one asks you to try to use the Go implementation of AES.

1. Create a Go package with methods KeyGen, Encrypt and Decrypt, that implement RSA key generation, encryption and decryption. Your solution should use integers from the math/big package.

    The KeyGen method should take as input an integer $k$, such that the bit length of the generated modulus $n=p q$ is $k$. The primes $p, q$ do not need to be primes with certainty, they only need to be "probable primes".

    The public exponent $e$ should be 3 (the smallest possible value, which gives the fastest possible encryption). This means that the primes $p, q$ that you output must be such that

    $$
    \operatorname{gcd}(3, p-1)=\operatorname{gcd}(3, q-1)=1 .
    $$

    Recall that $e=3$ and $d$ must satisfy that $3 d \bmod (p-1)(q-1)=1$. Another way to express this is to say that $d$ must be the inverse of 3 modulo $(p-1)(q-1)$, this is written

    $$
    d=3^{-1} \bmod (p-1)(q-1) .
    $$

    This way to express the condition will be useful when computing $d$.
    Facts you may find useful:

   - Other than standard methods for addition and multiplication, Mod and ModInverse will be useful.
   - To generate cryptographically secure randomness, use crypto/rand. In particular, the function Prime from the crypto/rand package may be helpful to you here.

    Test your solution by verifying (at least) that your modulus has the required length and that encryption followed by decryption of a few random plaintexts outputs the original plaintexts. Note that plaintexts and ciphtertexts in RSA are basically numbers in a certain interval. So it is sufficient to test if encryption of a number followed by decryption returns the original number. You do not need to, for instance, convert character strings to numbers.

2. Implement methods ``EncryptToFile`` and ``DecryptFromFile`` that encrypt and decrypt using AES in counter mode, using a key that is supplied as input. The ``EncryptToFile`` method should take as input a file name and should write the ciphertext to the file. Conversely the ``DecryptFromFile`` method should read the ciphertext from the file specified, decrypt and output the plaintext. Test your solution by encrypting a secret RSA key to a file. Then decrypt from the file, and check that the result can be used for RSA decryption.

    Go?s official crypto package has an AES implementation, so importing "crypto/aes? does the trick. The key-size of the key provided to the ``aes.NewCipher ([] byte(key goes here))`` call determines the strength of the cipher, so make sure its 16, 24, or 32 bytes, otherwise it will error (this is documented behaviour, but reiterating just in case).

### Besvarelse til 5.11

## Exercise 6.10 (RSA signatures)

### Opgave beskrivelse til 6.10

Extend your Go package from Exercise 5.11 so that it can generate and verify RSA signatures, where the message is first hashed with SHA-256 and then the hash value is signed using RSA, as described in Sec. 6.4. The hashing can be done with the ``crypto/sha256`` package.

Note: international standards for signatures always demand that the hash value is padded in some way before being passed to RSA, but you can ignore this here. Thus the hash value (which will be returned as a byte array) can be converted to an integer directly. Such direct conversion should not be done in a real application.
In addition to the code, your solution should contain the following:

1. Verify that you can both generate and verify the signature on at least one message. Also modify the message and check that your verification rejects.

2. Measure the speed at which you can hash, in bits per second. For this you should time the hashing of messages much longer than a hash value, in order to get realistic data say $10 \mathrm{~KB}$;

3. Measure the time you code spends to produce an RSA signature on a hash value when using a 2000 bit RSA key;

4. Assume you had to process the entire message using RSA. Use the result from question 3 to compute the speed at which you could do this (in bits per second). Hint: one of the RSA operations you timed in question 3 would allow you to process about 2000 bits. Compare your result to the speed you measured in question 2. Does it look like hashing makes signing more efficient?

### Besvarelse til 6.10
