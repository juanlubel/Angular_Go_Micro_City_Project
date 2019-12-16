import {initialUserState, IUserState} from '../state/user.state';
import {EUserActions, UserActions} from '../actions/user.actions';


export const userReducers = (
  state = initialUserState,
  action: UserActions
): IUserState => {
  switch (action.type) {
    case EUserActions.GetUsersSuccess: {
      return {
        ...state,
        users: action.payload
      };
    }
    case EUserActions.GetUserSuccess: {
      return {
        ...state,
        selectedUser: action.payload
      };
    }
    case EUserActions.UpdateUserSuccess: {
      return {
        ...state,
        users: action.payload
      };
    }
    // case EUserActions.GetTokenSuccess: {
    //   return {
    //     ...state,
    //     token: action.payload
    //   };
    // }

    default:
      return state;
  }
};
