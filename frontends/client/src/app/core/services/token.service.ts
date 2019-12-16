import {Injectable} from '@angular/core';
import {ApiService} from './api.service';
import {TokenType} from '../../models';

@Injectable(
  {providedIn: 'root'}
)
export class TokenService {
  public constructor(
    private apiService: ApiService
  ) {
  }
  public getToken(payload: TokenType) {
    const type = payload.type;
    const user = payload.name;
    console.log('get token', type, user);
    return this.apiService.get('redis', type, `/${user}`).pipe();
  }

}
