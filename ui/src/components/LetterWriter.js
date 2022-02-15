import React from "react";
import { useInput } from './hooks/input-hook';
import '../styles/LetterWriter.css';

export function LetterWriter(props) {
  const { value:author, bind:bindAuthor, reset:resetAuthor } = useInput('');
  const { value:recipient, bind:bindRecipient, reset:resetRecipient } = useInput('');
  const { value:heading, bind:bindHeading, reset:resetHeading } = useInput('');
  const { value:message, bind:bindMessage, reset:resetMessage } = useInput('');
  
  const handleSubmit = async(evt) => {
      evt.preventDefault();
      try {
        let res = await fetch("http://localhost:8080/api/letters", {
          method: "POST",
          body: JSON.stringify({
            FromName: author,
            ToName: recipient,
            Heading: heading,
            Body: message
          })
        })
        let resJson = await res.json();
        if (res.status !== 201) {
          console.error("Error sending letter")
        }
      } catch (err) {
        console.log(err);
      }

      resetRecipient();
      resetHeading();
      resetMessage();
  };

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