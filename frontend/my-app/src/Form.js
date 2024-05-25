import React, { useState } from "react";

function Form() {
  const [prefix, setPrefix] = useState("");
  const [suffix, setSuffix] = useState("");
  const [substring, setSubstring] = useState("");
  const [grade, setGrade] = useState("");
  const [type, setType] = useState("");
  const [result, setResult] = useState(null);


  const handleSubmit = async (event) => {
    event.preventDefault();

    const requestBody = {
      suffix: suffix,
      prefix: prefix,
      substring: substring,
      filter_by: {
        grade: grade,
        type: type,

      },
    };

    console.log('Request body:', requestBody);

    try {
      const response = await fetch("http://localhost:8080/search", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestBody)
      });

      if(response.ok) {
        const jsonResponse = await response.json();
        setResult(jsonResponse);
      } else {
        setResult(response.status + " " + response.statusText)
        console.log("HTTP-Error: " + response.status);
      }
    } catch (error) {
      console.error("Error: ", error.error);
    }
  };

  return (
      <form onSubmit={handleSubmit}>
        <label>
          Msidsn Prefix:
          <input type="text" value={prefix} onChange={e => setPrefix(e.target.value)} />
        </label>
        <label>
          Msidsn Suffix:
          <input type="text" value={suffix} onChange={e => setSuffix(e.target.value)} />
        </label>
        <label>
          Msidsn Substring:
          <input type="text" value={substring} onChange={e => setSubstring(e.target.value)} />
        </label>
        <label>
          Grade:
          <input type="text" value={grade} onChange={e => setGrade(e.target.value)} />
        </label>
        <label>
          Type:
          <input type="text" value={type} onChange={e => setType(e.target.value)} />
        </label>
        <button type="submit">Search</button>
        {result && <pre>{JSON.stringify(result, null, 2)}</pre>}
      </form>
  );
}

export default Form;