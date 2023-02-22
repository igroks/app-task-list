import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SliderToggleComponent } from './slider-toggle.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [SliderToggleComponent],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule
  ],
  exports: [SliderToggleComponent]
})

export class SliderToggleModule { }
