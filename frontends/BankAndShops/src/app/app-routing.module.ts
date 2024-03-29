import { NgModule } from '@angular/core';
import { Routes, RouterModule, PreloadAllModules } from '@angular/router';

const routes: Routes = [
  {
    path: "inicio",
    loadChildren: "./home/home.module#HomeModule"
  },
  {
    path:"", pathMatch:"full" ,redirectTo:"/inicio"
  },
  {
    path: "banks",
    loadChildren: "./bank/bank.module#BankModule"
  },
  {
    path: "forum",
    loadChildren: "./forum/#ForumModule"
  }
  /* {
    path: "king",
    loadChildren: "../app/king/king.module#KingModule"
  } */
];
@NgModule({
  imports: [RouterModule.forRoot(routes, {
    // preload all modules; optionally we could
    // implement a custom preloading strategy for just some
    // of the modules (PRs welcome 😉)
    /* preloadingStrategy: PreloadAllModules */
  })],
  exports: [RouterModule]
})
export class AppRoutingModule {}