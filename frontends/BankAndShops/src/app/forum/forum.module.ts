import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule} from '@angular/forms';
import { MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
  MatFormFieldModule,MatInputModule,MatGridListModule,MatDividerModule,MatCardModule } from  '@angular/material';

import { ForumRoutingModule } from './forum-routing.module';
import { ForumComponent } from './forum.component';
import { TopicComponent } from './topics/topic.component';
import { CommentComponent } from './comments/comment.component';
import { TopicResolver } from './topics/forum.resolver';
import { CommentResolver } from './comments/comments.resolve';


@NgModule({
  declarations: [ForumComponent,TopicComponent,CommentComponent],
  imports: [CommonModule,FormsModule, ReactiveFormsModule,
    ForumRoutingModule,
    MatToolbarModule, MatIconModule, MatSidenavModule, MatListModule, MatButtonModule,
    MatFormFieldModule,MatInputModule,MatGridListModule,MatDividerModule,MatCardModule],
  providers: [
    TopicResolver,CommentResolver
  ],
  exports:[]
})
export class ForumModule {}