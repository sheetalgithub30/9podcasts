import { Button, Row, Col } from "antd";
import React from "react";

const NoPodcast = () => {
  return (
    <Row justify="center" align="middle" style={{ height: "500px" }}>
      <Col>
        <p>You do not have any podcasts.</p>
        <Button type="primary">Create a new podcast</Button>
      </Col>
    </Row>
  );
};

export default NoPodcast;
