import { Injectable } from '@angular/core';

import { Resolve, Router, ActivatedRouteSnapshot, RouterStateSnapshot, ActivatedRoute } from '@angular/router';

import { Topic,ForumService } from 'src/app/core';

import { Observable } from 'rxjs';

import { catchError } from 'rxjs/operators';

@Injectable()
export class CommentResolver implements Resolve<Topic> {
  constructor(
    private router: Router,
    private ForumService: ForumService,
    private route: ActivatedRoute
  ) {}

  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> {
    console.log(route.params.topic_name)
    return this.ForumService.getTopic(route.params.topic_name).pipe(catchError((err) => this.router.navigateByUrl('/')));

  }
}