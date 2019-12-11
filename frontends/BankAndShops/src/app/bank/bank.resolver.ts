import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';

import { BankService, Banks } from '../core';
import { take, catchError } from 'rxjs/operators';

@Injectable()
export class BankResolver implements Resolve<Banks> {
  constructor(
    private router: Router,
    private BankService: BankService
  ) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {

    return this.BankService.getAll().pipe(catchError((err) => this.router.navigateByUrl('/')));

  }
}