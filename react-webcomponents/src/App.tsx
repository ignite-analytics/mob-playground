import { useState, useRef, useEffect, HTMLAttributes, MutableRefObject, HtmlHTMLAttributes, RefObject  } from "react"
import "./App.css"
import { Button as CountingButton } from "./CountingButton"
import '@vaadin/select'
import {Select} from '@vaadin/select'
import '@vaadin/button'
import { Button } from '@vaadin/button'

// Also an option:
// https://github.com/lit/lit/tree/main/packages/labs/react

declare global {
  namespace JSX {

    interface VaadinSelectAttrs extends HTMLAttributes<HTMLElement> {
      items: any;
      label: string;
      value?: string;
      ref: RefObject<Select>;
    }

    interface VaadinButtonAttrs extends HTMLAttributes<HTMLElement> {
      theme: string
      ref: RefObject<Button>;
    }
    
    interface IntrinsicElements {
      'vaadin-select': VaadinSelectAttrs;
      'vaadin-button': any
    }
  }
}

function App() {
  const vaadinSelect = useRef<Select>(null)
  const vaadinButton = useRef<Button>(null)
  const items = [{
    label: 'aaa',
    value: '3'
  },
  {
    label: 'bbb',
    value: '4'
  },]


  useEffect(() => {
    if (vaadinSelect.current) {
      vaadinSelect.current.items = items;
      vaadinSelect.current.onchange = () => alert(vaadinSelect.current?.value)
    }
    if(vaadinButton.current) {
      vaadinButton.current.onclick = () => alert("hi from button")
    }
  })

  const strings = ["1", "2"];

  const val = '3';

  return (
    <div className="App">
      <CountingButton/>

      <vaadin-button ref={vaadinButton} theme="primary">Click me</vaadin-button>

      <vaadin-select
        ref={vaadinSelect}
        items={strings}
        value={val}
        label="Sort by"
      ></vaadin-select>
    </div>
  );
}

export default App;
