import { reactive } from "vue";

export default {
  state: reactive({
    url: "",
    path: "",
    query: "",
    hash: "",
    hashMap: {}, // 经过字典索引
    queryMap: {}, // 经过字典索引
  }),
};
