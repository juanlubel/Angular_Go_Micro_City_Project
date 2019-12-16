import {User} from '../../models';


export interface IUserState {
  users: User[];
  selectedUser: User;
  token: string;
}

export const initialUserState: IUserState = {
  users: null,
  selectedUser: null,
  token: null
}
