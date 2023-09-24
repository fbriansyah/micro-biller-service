
# Micro Biller Service

Aplikasi ini merupakan micro service untuk menyediakan endpoint pembayaran. Project berikut ini bertujuan untuk membuat proof of concept membagi sistem pembayaran menjadi beberapa micro service. Terdapat beberapa endpoint yang terdapat pada aplikasi ini, yaitu:

1. Inquiry
2. Payment
3. Advice (Check payment status).

## Inquiry
Inquiry merupakan endpoint untuk mengetahui tagihan dari nomor billing dari pelanggan. Berikut ini merupakan data yang dikeluarkan oleh endpoint:
1. Request mode (mode: INQ)
1. Response code (code)
1. Response message (message)
1. Nomor billing (bill_number)
1. Nama (name)
1. Billing amount (base_amount)
1. Denda (fine_amount)
1. Total Tagihan (total_amount = base_amount + fine_amount)

Data diatas hanya sebagian data yang biasanya dikeluarkan. Berikut ini merupakan spec teknis secara detail:
```
url: /inquiry
method: POST
header:
    - Authorization: bearier
    - Content-Type: application/json
body:
    - bill_number
```

## Payment
Payment merupakan edpoint untuk melakukan pelunasan tagihan dari nomor billing pelanggan. Berikut ini merupakan data yang dikeluarkan oleh endpoint:
1. Request mode (mode:PAY)
1. Response code (code)
1. Response message (message)
1. Nomor billing (bill_number)
1. Nama (name)
1. Total Tagihan (total_amount = base_amount + fine_amount)
Berikut ini merupakan spec teknis secara detail:
```
url: /payment
method: POST
header:
    - Authorization: bearier
    - Content-Type: application/json
body:
    - bill_number
    - total_amount
    - refference_number
```

## Advice
Advice merupakan endpoint untuk melakukan checking dari status tagihan. Ada kalanya ketika pembayaran terjadi kendala jaringan sehingga mendapat error timeout, karena pada dasar Biller Service ini berada di server yang berbeda dari cluster payment micro service.

Sehingga tujuan dari Endpoint ini untuk mengetahui apakah tagihan sebelumnya lunas atau gagal. Berikut ini merupakan data yang dikeluarkan oleh endpoint:
1. Request mode (mode: ADV)
1. Response code (code)
1. Response message (message)
1. Nomor billing (bill_number)
1. Nama (name)
1. Total Tagihan (total_amount = base_amount + fine_amount)
Berikut ini merupakan spec teknis secara detail:
```
url: /advice
method: POST
header:
    - Authorization: bearier
    - Content-Type: application/json
body:
    - bill_number
    - total_amount
    - refference_number
```