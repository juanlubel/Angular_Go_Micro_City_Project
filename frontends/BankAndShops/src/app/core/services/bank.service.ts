import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';

import { ApiService } from './api.service';
import { map, catchError } from 'rxjs/operators';
import { Banks } from '../models';

@Injectable()
export class BankService {
  constructor (
    private apiService: ApiService
  ) {}
  private formatErrors(error: any) {
    return  throwError(error.error);
  }

  getAll(): Observable<[Banks]> {
    return this.apiService.get('banks','api_url')
          .pipe(map(data => data.banks));
  }
  addBank (bank:Banks): Observable<Banks> {
    return this.apiService.post('banks','api_url', bank)
      .pipe(
        catchError(this.formatErrors)
      );
  }
  login_bank (bank:Banks): Observable<any> {
    
    return this.apiService.post('account/login','api_url', bank)
      .pipe(
        catchError(this.formatErrors)
      );
  }
}