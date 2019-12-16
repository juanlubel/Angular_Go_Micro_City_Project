export interface User {
  'name': string;
  'email': string;
  'slug': string;
  'nCard': string;
}

export interface IUserHttp {
  'users': User[];
}

export interface TokenType {
  'name': string;
  'type': string;
}

export interface TokenHttp {
  'token': string;
}
