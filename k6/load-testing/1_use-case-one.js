import http from "k6/http";
import { sleep } from "k6";

export const options = {
  vus: 25,
  duration: "600s",
};
const url = "http://34.102.130.42";

const products = [
  "0PUK6V6EV0",
  "1YMWWN1N4O",
  "2ZYFJ3GM2N",
  "66VCHSJNUP",
  "6E92ZMYYFZ",
  "9SIQT8TOJO",
  "L9ECAV7KIM",
  "LS4PSXUNUM",
  "OLJCESPC7Z",
];

const currencies = ["EUR", "USD", "JPY", "GBP", "TRY", "CAD"];

export default function () {
  index();
  setCurrency();
  browseProduct();
  addToCart();
  viewCart();
  checkout();
  sleep(randomise(0,10));
}

function index() {
  http.get(url);
}

function setCurrency() {
  let data = { currency_code: currencies[randomise(0,3)] };

  http.post(url + "/setCurrency", JSON.stringify(data));
}

function browseProduct() {
  http.get(url + "/product/" + products[randomise(0,8)]);
}

function addToCart() {
  let product = products[randomise(0,8)];

  let data = { product_id: product, quantity: randomise(1,5) };

  http.get(url + "/product/" + product);
  http.post(url + "/cart", data);
}

function viewCart() {
  http.get(url + "/cart");
}

function checkout() {
  addToCart();

  let data = {
    email: "master@stud.fh-campuswien.ac.at",
    street_address: "226 Favoritenstraße",
    zip_code: "94043",
    city: "Mountain",
    state: "CA",
    country: "United States",
    credit_card_number: "4432-8015-6152-0454",
    credit_card_expiration_month: "7",
    credit_card_expiration_year: "2023",
    credit_card_cvv: "671",
  };
  http.post(url + "/cart/checkout", data);
}

function randomise(min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min) + min);
}

