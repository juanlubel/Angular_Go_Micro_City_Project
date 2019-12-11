export interface User {
  'name': string;
  'email': string;
  'slug': string;
  'nCard': string;
}

export interface IUserHttp {
  'users': User[];
}
