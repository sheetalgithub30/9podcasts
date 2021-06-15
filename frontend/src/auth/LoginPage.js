import { Col, Row } from "antd";
import React from "react";
import Login from "./Login";

const LoginPage = () => {
  return (
    <Row align="middle" justify="center" style={{ height: "100vh" }}>
      <Col span={8}>
        <Login />
      </Col>
    </Row>
  );
};

export default LoginPage;
