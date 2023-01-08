import http from "k6/http";
import { sleep } from "k6";

export const options = {
  vus: 100,
  duration: "600s",
};

const url = "http://34.90.230.154";

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
  sleep(randomise(10));
}

function index() {
  http.get(url);
}

function setCurrency() {
  let data = { currency_code: currencies[randomise(3)] };

  http.post(url + "/setCurrency", JSON.stringify(data));
}

function browseProduct() {
  http.get(url + "/product/" + products[randomise(8)]);
}

function addToCart() {
  let product = products[randomise(8)];

  let data = { product_id: product, quantity: randomise(10) };

  http.get(url + "/product/" + product);
  http.post(url + "/cart", JSON.stringify(data));
}

function viewCart() {
  http.get(url + "/cart");
}

function checkout() {
  addToCart();

  let data = {
    email: "master@stud.fh-campuswien.ac.at",
    street_address: "Favoritenstraße 226",
    zip_code: "94043",
    city: "Vienna",
    state: "Vienna",
    country: "Austria",
    credit_card_number: "2332-2015-6112-2374",
    credit_card_expiration_month: "7",
    credit_card_expiration_year: "2023",
    credit_card_cvv: "183",
  };

  http.post(url + "/cart/checkout", JSON.stringify(data));
}

function randomise(parameter) {
  return Math.random() * parameter;
}
