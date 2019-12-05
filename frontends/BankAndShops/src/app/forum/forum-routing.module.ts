import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import {ForumComponent} from "./forum.component"



const routes: Routes = [
    {
        path: '',
        component: ForumComponent
        /* resolve: {
          banks: BankResolver
        } */
      }

];

@NgModule({
    exports: [RouterModule],
    imports: [RouterModule.forChild(routes)]
})
export class ForumRoutingModule { }