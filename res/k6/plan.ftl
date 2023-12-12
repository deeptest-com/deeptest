import http from "k6/http";

export const options = {
  scenarios: {
    ${scenarios}
  },
};

export default function () {
  ${codes}
}