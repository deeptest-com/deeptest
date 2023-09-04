import {WsMsgCategory} from "@/utils/enum";

export interface WsMsg {
  msg: string
  category: WsMsgCategory,
  data: any
  time: string

  token: string
  tokenType: string

  conn: string
}

export interface QueryParams {
  keywords: string,
  enabled: string,
  page: number,
  pageSize: number,
}

export interface QueryResult {
  result: any[];
  pagination: PaginationConfig;
}

export interface PaginationConfig {
  total: number;
  page: number;
  pageSize: number;
  showSizeChanger: boolean;
  showQuickJumper: boolean;
}



