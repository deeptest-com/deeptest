export interface InterfaceDetail {
  contentType: string;
  contentLang: string;
  headers: any[];
  cookies: any[];
  [key: string]: any;
}