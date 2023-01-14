import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '2m', target: 100 }, // below normal load
    { duration: '2m', target: 200 }, // normal load
    { duration: '2m', target: 300 }, // around the breaking point
    { duration: '2m', target: 400 }, // beyond the breaking point
    { duration: '2m', target: 500 },
  ],
};

export default function () {
  const BASE_URL = 'http://localhost:8080'; 

  http.batch([
    ['GET', `${BASE_URL}/product/OLJCESPC7Z`],
    ['GET', `${BASE_URL}/product/66VCHSJNUP`],
    ['GET', `${BASE_URL}/product/2ZYFJ3GM2N`],
    ['GET', `${BASE_URL}/product/L9ECAV7KIM`],
  ]);

  sleep(1);
}