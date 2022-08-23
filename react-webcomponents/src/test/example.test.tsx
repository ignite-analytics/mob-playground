import { render, fireEvent, screen } from '@testing-library/react';
import App from '../App'


test('123', ()=>{
    render(<App/>);

    expect(screen.getByRole('button')).toBeTruthy();
});
