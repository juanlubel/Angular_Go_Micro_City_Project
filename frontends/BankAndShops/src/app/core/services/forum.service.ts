import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
  
import { ApiService } from './api.service';
import { Topic, TopicWithComments } from '..';
import { map, catchError } from 'rxjs/operators';
import { Comment } from '../models';

  @Injectable()
  export class ForumService {
    constructor (
      private apiService: ApiService
    ) {}
    private formatErrors(error: any) {
      return  throwError(error);
    }
    getAllTopics(): Observable<[Topic]> {
      return this.apiService.get('topics',"prequery_url")
            .pipe(map(data => data.topics));
    }
    createTopic (topic:Topic): Observable<Topic> {
      console.log(topic)
      return this.apiService.post('topics','forum_url', topic)
        .pipe(
          catchError(this.formatErrors)
        );
    }
    createComment (comment:Comment,Topic:String): Observable<Comment> {
      console.log(comment)
      return this.apiService.post('topic/comment','forum_url', comment)
        .pipe(
          catchError(this.formatErrors)
        );
    }
    deleteTopic (topic:string): Observable<Topic> {
      console.log(topic)
      return this.apiService.delete('topic/'+topic,'forum_url')
        .pipe(
          catchError(this.formatErrors)
        );
    }
    getTopic(topic_title:string):Observable<TopicWithComments>{
      return this.apiService.get(`topics/`+topic_title,"forum_url")
      .pipe(map(data => data.topic));
    }
  }