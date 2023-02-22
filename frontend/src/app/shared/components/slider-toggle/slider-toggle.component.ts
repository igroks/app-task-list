import { Component, Input } from '@angular/core';
import { FormControl } from '@angular/forms';

@Component({
  selector: 'app-slider-toggle',
  templateUrl: './slider-toggle.component.html',
  styleUrls: ['./slider-toggle.component.scss']
})

export class SliderToggleComponent {
  @Input() control!: FormControl<any>;
}
