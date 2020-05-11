[![Build Status](https://travis-ci.com/Hermsi1337/go-magento2.svg?branch=master)](https://travis-ci.com/Hermsi1337/go-magento2)
[![Maintainability](https://api.codeclimate.com/v1/badges/f9e75064d22478ed207f/maintainability)](https://codeclimate.com/github/Hermsi1337/go-magento2/maintainability)
[![Donate](https://img.shields.io/badge/Donate-PayPal-yellow.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=T85UYT37P3YNJ&source=url)

# go-magento2
A Golang package for communicating with the magento2 api. (tested with >=2.1.0)
   
I initially built this package because I need it for a project I'm currently working on.  
For the start, I will add further features upon my need.

If you need a feature which is not implemented yet, feel free to open a pull request.  

## Features
* [x] Guest api
  * [x] guest-carts
    - [x] add items
    - [x] get available shipping carrier
    - [x] add shipping information (billing- and shipping-address)
    - [x] get available payment methods
    - [x] add payment method
    - [x] place order
* [x] Registered customer api
  * [x] cart
    - [x] add items
    - [x] get available shipping carrier
    - [x] add shipping information (billing- and shipping-address)
    - [x] get available payment methods
    - [x] add payment method
    - [x] place order
* [x] Administrator / Integration
  * [x] cart
    - [x] all features from guest- and customer-api
  * [x] order
    - [x] add comment to order
  * [x] products
    - [x] create new product
        - [x] configurable products
            - [x] assign simple products
        - [x] simple products
    - [x] create attributes
    - [x] create attribute-set
      - [x] assign attributes to attribute-set

## Examples
See [examples-directory](examples/).
