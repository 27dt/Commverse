import { useState } from "react";
import "../styles/Chat.css";

function Chat() {
	const [messages, setMessages] = useState([])

	function sendMessage(formData: FormData) {
		console.log(formData.get("message"));
	}
	
	return (
		<div>
			<form action={sendMessage}>
				<input name="message" />
				<button type="submit">Send Message</button>
			</form>
		</div>
	);
}

export default Chat;
