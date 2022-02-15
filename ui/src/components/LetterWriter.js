import React from "react";
import { useInput } from './hooks/input-hook';
import './LetterWriter.css';

export function LetterWriter(props) {
  const { value:author, bind:bindAuthor, reset:resetAuthor } = useInput('');
  const { value:recipient, bind:bindRecipient, reset:resetRecipient } = useInput('');
  const { value:heading, bind:bindHeading, reset:resetHeading } = useInput('');
  const { value:message, bind:bindMessage, reset:resetMessage } = useInput('');
  
  const handleSubmit = (evt) => {
      evt.preventDefault();
      alert(`Submitting Message`);
      resetRecipient();
      resetHeading();
      resetMessage();
  }
  return (
    <form onSubmit={handleSubmit} className="LetterWriter">
      <label>
        Your Name:
        <input type="text" {...bindAuthor} />
      </label>
      <label>
        Recipient
        <input type="text" {...bindRecipient} />
      </label>
      <label>
        Heading
        <input type="text" {...bindHeading} />
      </label>
      <label>
        Message
        <input type="text" {...bindMessage} />
      </label>
      <input type="submit" value="Submit" />
    </form>
  );
}