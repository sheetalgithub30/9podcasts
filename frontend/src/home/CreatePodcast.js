import { LeftOutlined } from "@ant-design/icons";
import { Form, Modal, Button, Row, Col , Input , Select} from "antd";
import React, { useState } from 'react';

const { Option } = Select;

const CreatePodcast = () => {
    const [isModalVisible, setIsModalVisible] = useState(false);

  const showModal = () => {
    setIsModalVisible(true);
  };

  const handleOk = () => {
    setIsModalVisible(false);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };
  const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 12 },
    
  };
  const validateMessages = {
    required: `Label is required!`
  };


    const onFinish = (values) => {
      console.log(values);
    };
  return (
    <Row justify="center" align="middle" style={{ height: "500px" }}>
      <Col>
        <p>You do not have any podcasts.</p>
        <Button type="primary" onClick={showModal} >Create a new podcast</Button>
        <Modal title="New Podcast" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel} 
        width={650} height={1000}>
        <Form {...layout} name="nest-messages" onFinish={onFinish}
         validateMessages={validateMessages} align={LeftOutlined}>
        <h3>Podcast Details</h3>
        <Form.Item name={['user', 'name']} label="Title" >
        <Input />
      </Form.Item>
      <Form.Item name={['user', 'name']} label="Description" >
        <Input.TextArea />
      </Form.Item>
      <Form.Item name={['user', 'name']} label="Public website and website feed address" >
      <Input />
      </Form.Item>
      <Form.Item label="Language">
        <Input.Group compact>
          <Form.Item
            name={'language'}
            noStyle
            rules={[{ required: true, message: 'Province is required' }]}
          >
            <Select placeholder="English">
              <Option value="English">English</Option>
              <Option value="Hindi">Hindi</Option>
            </Select>
            </Form.Item>
            </Input.Group>  
            </Form.Item>  
            <p align='center'>This episode includes explicit content.</p>
        </Form>
      </Modal>
      </Col>
    </Row>
  );
};

export default CreatePodcast;
