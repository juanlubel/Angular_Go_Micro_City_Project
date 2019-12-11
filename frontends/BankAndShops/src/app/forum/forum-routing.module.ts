import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";


import {ForumComponent} from "./forum.component"
import { CommentComponent } from './comments/comment.component';
import { TopicResolver } from './topics/forum.resolver';
import { CommentResolver } from './comments/comments.resolve';

const routes: Routes = [
    {
        path: '',
        component: ForumComponent,
        resolve: {
          topics: TopicResolver
        }
      },
      {
        path: 'topic/:topic_name',
        component: CommentComponent,
        resolve: {
          topic: CommentResolver
        }
      }

];

@NgModule({
    exports: [RouterModule],
    imports: [RouterModule.forChild(routes)]
})
export class ForumRoutingModule { }